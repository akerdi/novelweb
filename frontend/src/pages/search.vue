<template lang="pug">
  .search-bg
    .containerView
      .flex-row-center.f-m-t-10
        el-input(v-model="search" placeholder="小说...")
        el-button.f-m-l-20(@click="searchFunc" type="primary") 搜索

      el-table.f-m-t-30(:data="novellist" style="width: 100%")
        el-table-column(prop="" label="序号" width="50px")
          template(slot-scope="scope")
            span {{scope.$index + 1}}
        el-table-column(prop="title" label="书名")
          template(slot-scope="scope")
            a(@click="chooseChapter(scope.row)") {{scope.row.title}}
        //- el-table-column(prop="url" label="链接")

</template>

<script>
import { search } from '@/service'
export default {
  name: 'search',
  data() {
    return {
      search: '',
      page: "1",
      novellist: [],
      loading: false
    }
  },
  methods: {
    searchFunc() {
      this.searchNovel()
    },
    async searchNovel() {
      const params = {
        p: this.page,
        q: this.search
      }
      this.loading = true
      const data = await search(params)
      this.loading = false
      this.novellist = data.data
      console.log("data::: ", data)
    },
    chooseChapter(row) {
      const query = {
        q: row.md5
      }
      console.log("%%%%%", row)
      this.$router.push({name: "Chapter", query})
    }
  },
  mounted() {
    if (this.$route.query) {
      this.search = this.$route.query.q
      this.searchNovel()
    }
  }
}
</script>

<style lang="scss" scoped>
  .search-bg {
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: cornflowerblue;
    .containerView {
      padding: 20px;
      margin-left: 20px;
      width: 50%;
      background-color: darkkhaki;
      min-width: 400px;
      el-button {
        height: 64px;
      }
    }
  }

</style>