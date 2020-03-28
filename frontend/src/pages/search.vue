<template lang="pug">
  .search-bg
    .containerView
      .flex-row-center.f-m-t-10
        el-input(v-model="searchText" @keyup.enter.native='handleSearch' placeholder="小说...")
        el-button.f-m-l-20(@click="handleSearch" type="primary") 搜索

      el-table.f-m-t-30(:data="novellist" style="width: 100%" v-loading="loading")
        el-table-column(prop="" label="序号" width="55px")
          template(slot-scope="scope")
            span {{scope.$index + 1}}
        el-table-column(prop="title" label="书名")
          template(slot-scope="scope")
            a(@click="chooseChapter(scope.row)")
              .flex-colume-center.chapterRow
                span {{scope.row.title}}
                span ({{scope.row.local ? "来源于[缓存]" : "来源于[匹配]"}})
      el-pagination.searchPagination(small @current-change="handleCurrentChange" :current-page.sync="currentPage" layout="prev, pager, next" :total="80" :page-size="10")
</template>

<script>
import { search } from '@/service'
export default {
  name: 'search',
  data() {
    return {
      searchText: '',
      novellist: [],
      loading: false,
      currentPage: 1
    }
  },
  methods: {
    handleSearch() {
      this.currentPage = 1
      this.search()
    },
    async search() {
      const params = {
        p: this.currentPage,
        q: this.searchText
      }
      this.loading = true
      const data = await search(params)
      this.loading = false
      const arrayLocal = data.data[0]
      const arrayNet = data.data[1]
      this.novellist = []
      for (const novelBook of arrayLocal) {
        novelBook.local = true
        this.novellist.push(novelBook)
      }
      for (const novelBook of arrayNet) {
        novelBook.local =false
        this.novellist.push(novelBook)
      }
    },
    chooseChapter(row) {
      const query = {
        q: row.md5
      }
      console.log("%%%%%", row)
      this.$router.push({name: "Chapter", query})
    },
    handleCurrentChange(value) {
      console.log(value, "   dddd ", this.currentPage)
      this.search()
    }
  },
  mounted() {
    if (this.$route.query) {
      this.searchText = this.$route.query.q
      this.handleSearch()
    }
  }
}
</script>

<style lang="scss" scoped>
  .search-bg {
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: #eeeeee;
    .containerView {
      padding: 20px;
      width: 50%;
      min-width: 330px;
      el-button {
        height: 64px;
      }
      .chapterRow {
        align-items: flex-start;
      }
      .searchPagination {
        margin-left: -5px;
        margin-top: 20px;
        margin-bottom: 20px;
        text-align: end;
      }
    }
  }

</style>