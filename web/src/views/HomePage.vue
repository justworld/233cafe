<template>
  <el-main class="main1 display1">
    <van-list v-model="loading" :finished="finished" finished-text="没有更多了" @load="load" class="main2" :offset="500">
      <div class="row1" v-for="i in articles" :key="i.id">
        <div class="row2">
          <el-link :underline="false" :href="'/info/'+i.id" class="link1 bold1" :key="i.id" target="_blank">{{i.title}}</el-link>
          <div class="cell1">{{i.desc}}</div>
          <div class="cell2">
            <van-icon name="eye" class="icon1" /><span class="font1">{{i.read}}</span>
            <span class="hover1" @click="like(i)"><van-icon name="like" class="icon1 color1"/><span class="font1">{{i.love}}</span></span>
          </div>
        </div>
        <div class="row3">
          <el-link :underline="false" :href="'/info/'+i.id" class="link1" :key="i.id" target="_blank"><img class="img1" :src="i.cover"></el-link>
        </div>
      </div>
    </van-list>
    <div class="main3 padding1 bgcolor1">
      <div class="padding2 border1">热门阅读</div>
      <div class="row1" v-for="i in hotSpots" :key="i.id">
        <div class="row4">
          <el-link :underline="false" :href="'/info/'+i.id" class="link1 link2" :key="i.id" target="_blank">{{i.title}}</el-link>
        </div>
        <div class="row5">
          <el-link :underline="false" :href="'/info/'+i.id" class="link1" :key="i.id" target="_blank"><img class="img2" :src="i.cover"></el-link>
        </div>
      </div>
    </div>
  </el-main>
</template>
<script type="text/ecmascript-6">
export default {
  data () {
    return {
      p: 1,
      s: '',
      loading: false,
      finished: false,
      articles: [],
      hotSpots: []
    }
  },
  created () {
    let $this = this
    $this.s = $this.$route.query.s
    $this.load()
    $this.loadHot()
  },
  methods: {
    load () {
      let $this = this
      if (!$this.loading) {
        $this.loading = true
        $this.$api.getArticles(res => {
          for (var i in res.data) {
            $this.articles.push(res.data[i])
          }
          $this.loading = false
          $this.finished = true
        }, { p: 1, cz: $this.$route.meta.cz, s: $this.s })
      }
    },
    loadHot () {
      let $this = this
      $this.$api.getHots(res => {
        $this.hotSpots = res.data
      })
    },
    like (i) {
      let $this = this
      $this.$api.addLike(res => {
        i.love = i.love + 1
      }, {id: i.id})
    }
  }
}
</script>
