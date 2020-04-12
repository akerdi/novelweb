<template lang="pug">
  .chapter-bd
    .novelTitle {{this.chapter.name || "无题"}}
    .chapterContainer(v-if="chapterTitleList.length")
      el-backtop(target=".chapterContainer" :visibility-height='150' :right="50" :bottom="50")
      .flex-row-center.novelContent(v-for="(list, i) in chapterTitleList")
        .flex-row-center.f-m-t-20(v-for="index in 4")
          a.f-m-r-10.chapterTitle(@click="chooseContent(list[index-1], i*4+index-1)") {{ list[index-1] }}
    .link_tips.f-m-b-20(v-if="this.chapter.OriginURL")
      span 该文章由网络获取, 如有侵权请联系QQ:767838865@qq.com 立即撤下.&nbsp
      a(:href="this.chapter.OriginURL" target="__blank") 原网址
</template>

<script>
import _ from 'lodash'
import callAsync from '@/lib/awaitCall'
import { chapter } from '@/service'
export default {
  name: "chapter",
  data() {
    return {
      md5: '',
      chapterList: [],
      title: '',
      chapterTitleList: [],
      chapter: {}
    }
  },
  methods: {
    async searchChapter() {
      const params = {
        md5: this.md5
      }
      const [err, res] = await callAsync(chapter(params))
      if (err) return this.$message.error(err.message)

      if (res.data) this.chapter = res.data.chapter
      if (!this.chapter) {
        return this.$message.error("没有找到数据")
      }
      const chapterList = this.chapter.Chapters
      const totalineNum = Math.ceil(chapterList.length/4)
      const novelList = []
      for (let i=0; i < totalineNum; i++) {
        const rowList = []
        for (let j=0; j < 4; j++) {
          const index = i*4+j
          const res = _.assign({}, chapterList[index])
          rowList.push(res.name)
        }
        novelList.push(rowList)
      }
      this.chapterTitleList = novelList
    },
    chooseContent(row, index) {
      const query = {
        q: this.md5,
        i: index
      }
      this.$router.push({name: "Content", query})
    }
  },
  mounted() {
    if (this.$route.query) {
      this.md5 = this.$route.query.q
      this.searchChapter()
    }
  }
}
</script>

<style lang="scss" scoped>
  .chapter-bd {
    .back-ball {
      background-color: tomato;
      color: #fff;
      border-radius: 22px;
      padding: 10px;
    }
    overflow: auto;
    .novelTitle {
      font-size: 18px;
      font-weight: bold;
      color: #555555;
      margin: 0 auto;
      margin-top: 20px;
      word-break:normal;
    }
    .chapterContainer {
      padding: 10px;
      margin-top: 10px;
      .novelContent {
        margin: 0 auto;
        margin-top: 0px;
        width: 99%;
        background-color: antiquewhite;
        justify-content: flex-start;
        .chapterTitle {
          justify-content: space-between;
          // padding: 4px;
          text-align: start;
          margin-left: 4px;
          margin-right: 4px;
        }
      }
    }

  }
</style>

