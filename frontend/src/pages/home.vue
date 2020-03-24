<template lang="pug">
  .home-bg
    .inputContainer
      .flex-colume-center.flexContainer
        img(src="../assets/images/logo.png")
        .flex-row-center.f-m-t-10
          el-input(v-model="search" placeholder="小说...")
          el-button.f-m-l-20(@click="searchFunc" type="primary") 搜索
</template>

<script>
export default {
  name: "home",
  data() {
    return {
      search: ''
    }
  },
  methods: {
    searchFunc() {
      console.log("home~~~~~~", this.search)
      if (_.isEmpty(this.search)) {
        return console.log("empty!!!")
      }
      const query = {q: this.search}
      this.$router.push({ name: "Search", query })
    }
  }
}
</script>

<style lang="scss" scoped>
  .home-bg {
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: antiquewhite;
    .inputContainer {
      border-radius: 25px;
      width: 40%;
      min-width: 500px;
      background-color: blueviolet;
      margin: 0 auto;
      margin-top: 20vh;
      .flexContainer {
        padding: 20px;
      }
      img {
        margin: 0 auto;
        width: 80px;
        height: 80px;
      }
    }

  }
</style>