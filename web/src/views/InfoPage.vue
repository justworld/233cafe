<template>
  <el-main class="main5">
    <div class="main3 main6 padding1 bgcolor1">
      <div class="padding2 border1">推荐阅读</div>
      <div class="row5" v-for="i in recommends" :key="i.id">
        <el-link :underline="false" :href="'/info/'+i.id" class="link1" :key="i.id" target="_blank"><img class="img3" :src="i.cover"></el-link>
        <el-link :underline="false" :href="'/info/'+i.id" class="link1 link2 margin1" :key="i.id" target="_blank">{{i.title}}</el-link>
      </div>
    </div>
    <div v-html="content" class="main4"></div>
    <div class="ad1 padding3 bgcolor1">
      <div class="padding2 border1">广告</div>
      <el-carousel :interval="5000" class="carousel1">
        <el-carousel-item v-for="i in ads" :key="i.url">
          <el-link :underline="false" :href="i.url" class="link1 link3" :key="i.url" target="_blank"><img class="img4" :src="i.cover"></el-link>
        </el-carousel-item>
      </el-carousel>
    </div>
  </el-main>
</template>

<script type="text/ecmascript-6">
export default {
  data () {
    return {
      title: '',
      content: '',
      recommends: [],
      ads: []
    }
  },
  created () {
    let $this = this
    $this.$api.getArticle(res => {
      $this.title = res.data.title
      document.title = $this.title
      $this.content = res.data.content
      $this.recommends = res.data.recommends
      $this.ads = res.data.ads
    }, { id: $this.$route.params.id })
  },
  methods: {
  }
}
</script>
