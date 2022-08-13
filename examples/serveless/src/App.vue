<template>
<div class="app-list">
  <p>剪贴板内容:  </p>

  <div v-for="(cc, id) in clip" :key="id">
    <h2>设备{{id}}:</h2>

    <pre>{{cc}}</pre>
  </div>
  <button @click="test_read">测试读取</button><p>测试结果:{{test_content}}</p>

</div>
</template>

<script>
import { lzcAPIGateway } from "@lazycatcloud/sdk"

export default {
  async mounted() {
    let cc = new lzcAPIGateway()

    window.cc = cc;
    let s = await cc.session
    console.log("当前信息:", s)

    this.devices = (await cc.devices.ListEndDevices({"uid": s.uid})).devices

    let dp = await cc.currentDevice
    window.dp = dp

    dp.clipboard.Watch({"mime":"text/plain"}).subscribe({
      next: x => {
        this.$set(this.clip, "local", new TextDecoder("utf-8").decode(x.content));
      },
      error: err => this.clip = err,
      complete: ()=> alert("clipboard terminated!"),
    })
  },
  methods: {
    async test_read() {
      let c = await window.dp.clipboard.Read({})
      console.log("CC:",c.content)
      this.test_content = new TextDecoder("utf-8").decode(c.content);
    }
  },
  data() {
    return {
      test_content:"",
      clip: {},
      "devices": [],
    };
  },
};
</script>
