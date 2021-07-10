<template>
  <el-main class="main5">
    <div class="main3 main6 padding1 bgcolor1">
      <div class="padding2 border1">推荐阅读</div>
      <div class="row5" v-for="i in recommends" :key="i.id">
        <el-link :underline="false" :href="'info/'+i.id" class="link1" :key="i.id" target="_blank"><img class="img3" :src="i.cover"></el-link>
        <el-link :underline="false" :href="'info/'+i.id" class="link1 link2 margin1" :key="i.id" target="_blank">{{i.title}}</el-link>
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
      recommends: [
        {id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 2, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/3a8435d2546cfd9a03b636b86ccbc8c2.jpg'},
        {id: 3, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 4, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/3a8435d2546cfd9a03b636b86ccbc8c2.jpg'}],
      ads: [{url: '', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'}, {url: '', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/3a8435d2546cfd9a03b636b86ccbc8c2.jpg'}, {url: '', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'}, {url: '', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'}]
    }
  },
  created () {
    let $this = this
    $this.$api.getArticle(res => {
      $this.title = res.data.title
      document.title = $this.title
      $this.content = res.data.content
    }, { id: $this.$route.params.id })
  },
  methods: {
  }
}
</script>
