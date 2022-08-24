<template>
  <div class="webdav">
    <div class="fm-title">
      <div class="fm-title-path">
        当前路径:
        <span>{{ pwd }}</span>
      </div>
    </div>

    <div class="fm-content">
      <table>
        <thead>
          <tr>
            <th>MIME类型</th>
            <th>最后更新</th>
            <th>大小</th>
            <th>是否文件夹</th>
            <th>文件名</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td></td>
            <td></td>
            <td></td>
            <td></td>
            <td class="fm-arrow" @click="handleLast">&#8629;..</td>
          </tr>
          <tr
            v-for="item in items"
            :key="item.filename"
            :title="`文件路径: ` + item.filename"
          >
            <td>{{ item.mime || "-" }}</td>
            <td>{{ item.lastmod }}</td>
            <td>{{ item.size }}</td>
            <td>{{ item.type === "directory" ? "是" : "否" }}</td>
            <td
              :class="[item.type === 'directory' ? 'is-dir' : 'is-file']"
              @click="handleFile(item)"
            >
              {{ item.basename }}
            </td>
            <td class="fm-delete" @click="handleDelete(item)">&#10007;</td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="fm-bottom"></div>
  </div>
</template>

<script>
import { createClient } from "webdav";

function pathJoin(path1, path2) {
  let res = path1 + "/" + path2;
  return (
    "/" +
    res
      .split("/")
      .filter((r) => r !== "")
      .join("/")
  );
}

export default {
  props: {
    webdavRoot: {
      type: String,
      required: true,
    },
  },
  async mounted() {
    this.fs = createClient(this.webdavRoot);
    this.enterDir(this.pwd);
  },
  data() {
    return {
      pwd: "/",
      items: [],
      fs: undefined,
    };
  },
  methods: {
    async isDir(path) {
      try {
        let res = await this.fs.stat(path);
        return res.type === "directory";
      } catch (e) {
        return Promise.reject(e);
      }
    },
    toPwd(path) {
      let p;
      if (path.startsWith("/")) {
        p = path;
      } else {
        p = pathJoin(this.pwd, path);
      }
      return p;
    },
    enterDir(path) {
      return this.fs
        .getDirectoryContents(path)
        .then((res) => {
          this.items = res;
          this.pwd = path;
        })
        .catch((err) => console.error(err));
    },
    download(path) {
      let res = this.fs.getFileDownloadLink(path);
      let basename = path.split("/").reverse()[0];
      let a = document.createElement("a");
      a.href = res;
      a.download = basename;
      a.click();
    },
    handleFile(item) {
      if (item.type === "directory") {
        return this.enterDir(this.toPwd(item.basename));
      } else {
        return this.download(this.toPwd(item.basename));
      }
    },
    handleLast() {
      let res = this.pwd.split("/").reverse().slice(1);
      this.enterDir("/" + res.reverse().join("/"));
    },
    handleDelete(item) {
      this.fs
        .deleteFile(item.filename)
        .then(() => {
          this.enterDir(this.pwd);
        })
        .catch((e) => {
          console.error(e);
        });
    },
  },
};
</script>

<style scoped>
.webdav {
  width: 100%;
}

.fm-title {
  margin: 14px;
}
.fm-title-path {
  display: flex;
  flex-direction: row;
}
.fm-title-path span {
  margin-left: 14px;
}

.fm-item {
  display: flex;
  flex-direction: row;
}
.fm-content {
  margin: 14px;
}
th,
td {
  padding: 0 14px;
}

.is-dir {
  color: #3daee9;
  cursor: pointer;
}
.is-file {
  color: #1cdc9a;
  cursor: pointer;
}
.is-dir:hover {
  text-decoration: underline;
}

.fm-arrow {
  color: #da4453;
}
.fm-delete {
  color: #da4453;
  cursor: pointer;
}
</style>
