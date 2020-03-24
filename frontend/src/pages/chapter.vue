<template lang="pug">
  .chapter-bd
    span aaaa
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
      console.log("#####", params)
      const res = await chapter(params)

      console.log("chapter$$$$$", res.data)
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

