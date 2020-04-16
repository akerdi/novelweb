<template lang="pug">
  .chapter-bd
    .f-p-10
      el-breadcrumb.f-m-t-20.f-m-b-20(separator-class="el-icon-arrow-right")
        el-breadcrumb-item(:to="{ path: '/' }") 首页
        el-breadcrumb-item(:to="{ path: '/search', query: {'q': novel} }") 搜索
        el-breadcrumb-item 章节列表
      .novelTitle {{this.chapter.name || "无题"}}
      .chapterContainer(v-if="chapterTitleList.length")
        el-backtop(target=".chapterContainer" :visibility-height='150' :right="50" :bottom="50")
        .flex-row-center.novelContent(v-for="(list, i) in chapterTitleList")
          .chapterRow(v-for="index in 4")
            a.chapterTitle(:href="list[index-1].href") {{ list[index-1].name }}
      .link_tips.f-m-t-20.f-m-b-20(v-if="this.chapter.OriginURL")
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
      chapter: {},
      novel: '' // 使用novel搜索词 进来的
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
      if (!this.chapter || !this.chapter.Chapters) {
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
          res.href = `/content?q=${this.md5}&i=${index}&n=${this.novel}`
          rowList.push(res)
        }
        novelList.push(rowList)
      }
      this.chapterTitleList = novelList
    },
    chooseContent(row, index) {
      const query = {
        q: this.md5,
        i: index,
        n: this.novel
      }
      this.$router.push({name: "Content", query})
    }
  },
  mounted() {
    if (this.$route.query) {
      this.md5 = this.$route.query.q
      this.novel = this.$route.query.n
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
      margin-top: 20px;
      .novelContent {
        margin: 0 auto;
        margin-top: 0px;
        width: 99%;
        background-color: antiquewhite;
        justify-content: space-around;
        .chapterRow {
          margin-top: 20px;
          width: 25%;
        }
        .chapterTitle {
          padding: 4px;
          text-align: center;
        }
      }
    }

  }
</style>

