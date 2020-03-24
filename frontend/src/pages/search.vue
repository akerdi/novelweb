<template lang="pug">
  .search-bg
    .containerView
      .flex-row-center.f-m-t-10
          el-input(v-model="search" placeholder="小说...")
          el-button.f-m-l-20(@click="searchFunc" type="primary") 搜索

</template>

<script>
export default {
  name: 'search',
  data() {
    return {
      search: ''
    }
  },
  methods: {
    searchFunc() {

    }
  },
  mounted() {
    if (this.$router.query) {
      this.search = this.$router.query.q
    }
    console.log("!!!!!", this.search)
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
    }
  }

</style>