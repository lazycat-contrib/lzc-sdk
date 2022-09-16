<template>
  <div class="app-list">
    <p>当前设备信息: {{ device_info }}</p>
    <p>剪贴板内容:</p>

    <div v-for="(cc, id) in clip" :key="id">
      <h2>设备{{ id }}:</h2>

      <pre>{{ cc }}</pre>
    </div>
    <button @click="test_read">测试读取</button>
    <p>测试结果:{{ test_content }}</p>

    <br />
    <WebDav v-if="webdavRoot" :webdavRoot="webdavRoot"></WebDav>
  </div>
</template>

<script>
import { lzcAPIGateway } from "@lazycatcloud/sdk";

async function getLocalWebdavRoot(cc) {
  let s = await cc.session;
  for (let d of (await cc.devices.ListEndDevices({ uid: s.uid })).devices) {
    if (s.deviceId == d.uniqueDeivceId) {
      return d.deviceApiUrl.replace(".d.", ".local.") + "/s/files/";
    }
  }
  return "";
}

export default {
  components: {
    WebDav: () => import("./webdav.vue"),
  },
  async mounted() {
    let cc = new lzcAPIGateway();

    window.cc = cc;
    let s = await cc.session;
    console.log("当前信息:", s);

    console.log("WEBDAV ROOT", await getLocalWebdavRoot(cc));
    this.webdavRoot = await getLocalWebdavRoot(cc);

    this.devices = (await cc.devices.ListEndDevices({ uid: s.uid })).devices;

    let dp = await cc.currentDevice;
    window.dp = dp;

    this.device_info = await dp.device.Query({});

    dp.clipboard.Watch({ mime: "text/plain" }).subscribe({
      next: (x) => {
        this.$set(
          this.clip,
          "local",
          new TextDecoder("utf-8").decode(x.content)
        );
      },
      error: (err) => (this.clip = err),
      complete: () => alert("clipboard terminated!"),
    });
  },
  methods: {
    async test_read() {
      let c = await window.dp.clipboard.Read({});
      console.log("CC:", c.content);
      this.test_content = new TextDecoder("utf-8").decode(c.content);
    },
  },
  data() {
    return {
      test_content: "",
      clip: {},
      devices: [],
      device_info: "loading",
      webdavRoot: "",
    };
  },
};
</script>
