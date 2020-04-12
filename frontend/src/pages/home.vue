<template lang="pug">
  .home-bg
    #inputContainer
      .flex-colume-center.flexContainer
        .flex-row-center.f-m-t-10
          el-input(v-model="search" @keyup.enter.native='handleSearch' placeholder="小说...")
          el-button.f-m-l-20(@click="handleSearch" type="primary") 搜索
    .recommand
      a.recommandTitle(v-for="(novel, i) in recommandNovels" @click="recommand(i)") {{novel}}

</template>

<script>
export default {
  name: "home",
  data() {
    return {
      search: '',
      recommandNovels: ["红楼梦", "神雕侠侣", "倚天屠龙记", "猛龙过江", "霸道人生", "神墓", "斗罗大陆"]
    }
  },
  methods: {
    handleSearch() {
      if (_.isEmpty(this.search)) {
        return console.log("empty!!!")
      }
      const query = {q: this.search}
      this.$router.push({ name: "Search", query })
    },
    recommand(i) {
      const novel = this.recommandNovels[i]
      this.search = novel
      this.handleSearch()
    }
  }
}
</script>

<style lang="scss" scoped>
  .home-bg {
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: #eeeeee;
    #inputContainer {
      background-image: url("../../static/muti2.jpg");
      background-repeat: repeat;
      background-position: 0 0;
      background-size: auto 100%;

      border-radius: 25px;
      width: 40%;
      min-width: 350px;
      margin: 0 auto;
      margin-top: 20vh;
      .flexContainer {
        padding: 20px;
      }
      // i
    }
    .recommand {
      width: 40%;
      min-width: 350px;
      white-space:pre-wrap;
      margin: 0 auto;
      margin-top: 20px;
      height: 150px;
      line-height: 48px;
      .recommandTitle {
        padding: 10px;
        margin-left: 10px;
        text-align: left;
        white-space:pre-wrap;
        margin-top: 20px;
      }
    }


  }
</style>