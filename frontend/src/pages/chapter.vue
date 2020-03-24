<template lang="pug">
  .chapter-bd
    el-table.f-m-t-30(:data="chapterList" style="width: 100%")
      el-table-column(prop="" label="序号" width="50px")
        template(slot-scope="scope")
          span {{scope.$index + 1}}
      el-table-column(prop="name" label="章节")
        template(slot-scope="scope")
          a(@click="chooseContent(scope)") {{scope.row.name}}
</template>

<script>
import { chapter } from '@/service'
export default {
  name: "chapter",
  data() {
    return {
      md5: '',
      chapterList: []
    }
  },
  methods: {
    async searchChapter() {
      const params = {
        md5: this.md5
      }
      const res = await chapter(params)
      this.chapterList = res.data.Chapters
    },
    chooseContent(scope) {
      const query = {
        q: this.md5,
        i: scope.$index
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

</style>

