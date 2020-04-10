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
      scrollTo(0,0);
    },
    handleKey(e) {
      switch (e.keyCode) {
        case 39:
          this.handleNext()
          break
        case 37:
        this.handleLast()
        break
      }
    },
    handleNext() {
      let index = this.index
      index++
      this.index = index
      const query = {
        q: this.md5,
        i: this.index
      }
      this.$router.replace({query})
    },
    handleLast() {
      let index = this.index
      index--
      this.index = index
      const query = {
        q: this.md5,
        i: this.index
      }
      this.$router.replace({query})
    },
    init() {
      if (this.$route.query) {
        const {q, i} = this.$route.query
        this.md5 = q
        this.index = i
        this.getContent()
      }
    }
  },
  beforeDestroy() {
    document.removeEventListener('keyup', this.handleKey)
  },
  mounted() {
    document.addEventListener('keyup', this.handleKey)
  },
  watch: {
    $route: {
      handler: 'init',
      immediate: true
    }
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