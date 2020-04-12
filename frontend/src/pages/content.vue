<template lang="pug">
  .content-bd(v-if="this.data")
    .novelTitle {{this.data.name || "无题"}}
    .elementTitle.f-m-t-10 {{this.data.element.name}}
    .flex-row-center.f-m-t-10
      el-button.f-m-r-30(type="text" @click="handleLast") 上一章
      el-button(type="text" size="medium" @click="handleNext") 下一章
    .f-m-t-20.f-m-b-20
      .contentHtml(v-html='this.data.content.content')
    .flex-row-center.f-m-t-10
      el-button.f-m-r-30(type="text" @click="handleLast") 上一章
      el-button(type="text" @click="handleNext") 下一章
    .link_tips.f-m-t-10(v-if="this.data")
      span 该文章由网络获取, 如有侵权请联系QQ:767838865@qq.com 立即撤下.

    i.loadingIcon(v-show="loading" class="el-icon-loading")
</template>

<script>
import { content } from '@/service'
export default {
  name: "contentvue",
  data() {
    return {
      md5: '',
      index: '',
      data: null,
      loading: false
    }
  },
  methods: {
    async getContent() {
      const params = {
        md5: this.md5,
        index: this.index
      }
      this.loading = true
      const res = await content(params)
      this.loading = false
      this.data = res.data
      scrollTo(0, 0);
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
    .elementTitle {
      font-size: 16px;
      color: #555555;
    }
    .loadingIcon {
      position: absolute;
      right: 48px;
      bottom: 48px;
      width: 50px;
      height: 50px;
    }
    .novelTitle {
      font-size: 18px;
      font-weight: bold;
      color: #555555;
      margin: 0 auto;
      margin-top: 40px;
    }
    .contentHtml {
      text-align: left;
      font-size: 17px;
      width: 70%;
      min-width: 300px;
      margin: 0 auto;
    }
  }
</style>