<template lang="pug">
  .content-bd
    .flex-row-center
      el-button(type="primary" @click="handleLast") Last
      el-button(type="primary" @click="handleNext") Next
    .f-m-t-20.f-m-b-20
      .contentHtml(v-html='html_content')
    .flex-row-center
      el-button(type="primary" @click="handleLast") Last
      el-button(type="primary" @click="handleNext") Next
</template>

<script>
import { content } from '@/service'
export default {
  name: "contentvue",
  data() {
    return {
      md5: '',
      index: '',
      html_content: ''
    }
  },
  methods: {
    async getContent() {
      const params = {
        md5: this.md5,
        index: this.index
      }
      const res = await content(params)
      this.html_content = res.data.content
    },
    handleKey(e) {
      let index = this.index
      switch (e.keyCode) {
        case 39: // right
          index = parseInt(this.index)
          index++
          break;
        case 37: // left
          index = parseInt(this.index)
          index--
          break
        default:
          return
      }
      // TODO 请求数据
      // location.href = `/content?q=${this.md5}&i=${index}`
      this.index = index
      this.getContent()
    },
    handleNext() {
      let index = this.index
      index++
      this.index = index
      this.getContent()
    },
    handleLast() {
      let index = this.index
      index--
      this.index = index
      this.getContent()
    },
  },
  beforeDestroy() {
    document.removeEventListener('keyup', this.handleKey)
  },
  mounted() {
    if (this.$route.query) {
      const {q, i} = this.$route.query
      this.md5 = q
      this.index = i
      this.getContent()
    }
    document.addEventListener('keyup', this.handleKey)
  }
}
</script>

<style lang="scss" scoped>
  .content-bd {
    padding: 20px 0px;
    .contentHtml {
      text-align: left;
      font-size: 17px;
      width: 70%;
      min-width: 300px;
      margin: 0 auto;
    }
  }
</style>