<template lang="pug">
  .chapter-bd
    .novelTitle {{title}}
    .f-m-t-10
      .flex-row-center.novelContent(v-for="(list, i) in chapterTitleList")
        .flex-row-center.f-m-t-20(v-for="index in 4")
          a.f-m-r-10.chapterTitle(@click="chooseContent(list[index-1], i*4+index-1)") {{ list[index-1] }}
</template>

<script>
import _ from 'lodash'
import { chapter } from '@/service'
export default {
  name: "chapter",
  data() {
    return {
      md5: '',
      chapterList: [],
      title: '',
      chapterTitleList: [],
    }
  },
  methods: {
    async searchChapter() {
      const params = {
        md5: this.md5
      }
      const res = await chapter(params)
      console.log("@@@", res)
      this.title = res.data.name || "无题"
      this.chapterList = res.data.Chapters
      const totalineNum = Math.ceil(this.chapterList.length/4)
      console.log("totalineNum: ", totalineNum)
      const novelList = []
      for (let i=0; i < totalineNum; i++) {
        const rowList = []
        for (let j=0; j < 4; j++) {
          const index = i*4+j
          const res = _.assign({}, this.chapterList[index])
          console.log("index:", index, "ree ", res)
          rowList.push(res.name)
        }
        novelList.push(rowList)
      }
      this.chapterTitleList = novelList
      console.log("n9vellist:: ", novelList)
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
    .novelTitle {
      font-size: 18px;
      font-weight: bold;
      color: #555555;
      margin: 0 auto;
      margin-top: 40px;
    }
    .novelContent {
      margin: 0 auto;
      margin-top: 0px;
      width: 70%;
      background-color: antiquewhite;
      justify-content: space-around;
      .chapterTitle {
        background-color: aquamarine;
        justify-content: space-between;
        padding: 8px;
      }
    }
  }
</style>

