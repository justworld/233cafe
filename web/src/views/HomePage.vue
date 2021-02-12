<template>
  <el-main class="main1 display1">
    <div class="main2">
      <div class="row1" v-for="i in articles" :key="i.id">
        <el-link :underline="false" :href="item.url" class="crumb-font2" :key="'link'+index">{{i.title}}</el-link>
      </div>
    </div>
    <div class="main3">
      <div class="">热门阅读</div>
      <div class="row2" v-for="i in hotSpots" :key="i.id">
        <div>{{i.title}}</div>
      </div>
    </div>
  </el-main>
</template>
<script type="text/ecmascript-6">
export default {
  data () {
    return {
      p: 1,
      categories: [{id: 1, title: '生活'}],
      articles: [],
      hotSpots: [
        {id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 2, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 3, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 4, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 5, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 6, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 7, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 8, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 9, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 10, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'}]
    }
  },
  created () {
    let $this = this
    $this.load()
  },
  methods: {
    load () {
      let $this = this
      let examples = [
        {id: 1, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10},
        {id: 2, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10},
        {id: 3, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10},
        {id: 4, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10},
        {id: 5, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10},
        {id: 6, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10},
        {id: 7, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10},
        {id: 8, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10},
        {id: 9, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10},
        {id: 10, category_id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', desc: '500多名救援人员正在争分夺秒地营救被困的34名工人（可能还不止）。这些工人被困在北阿坎德邦查莫里地区一水电工程1.6公里长的隧道里，隧道里满是碎石。...', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg', read: 10, love: 10}]
      for (var i in examples) {
        $this.articles.push(examples[i])
      }
    }
  }
}
</script>
