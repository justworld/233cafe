<template>
  <el-main class="main5">
    <div class="main3 main6">
      <div class="">推荐阅读</div>
      <div class="row5" v-for="i in recommends" :key="i.id">
        <el-link :underline="false" :href="'info/'+i.id" class="link1" :key="i.id" target="_blank"><img class="img3" :src="i.cover"></el-link>
        <el-link :underline="false" :href="'info/'+i.id" class="link1 link2" :key="i.id" target="_blank">{{i.title}}</el-link>
      </div>
    </div>
    <div v-html="info.content" class="main4"></div>
    <div class="ad1">
      <div class="">广告</div>
      <el-carousel :interval="5000">
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
      info: { title: '', content: '<p>转眼间，又到了一年的中秋。</p><p>早上起床吃饭，在家好好休息休息。</p><figure><noscript><img src="https://pic1.zhimg.com/v2-54f656619f5520e303cd5b56e3fee8d8_b.jpg" class="content_image"/></noscript><img src="data:image/svg+xml;utf8,&lt;svg xmlns=&#39;http://www.w3.org/2000/svg&#39; width=&#39;0&#39; height=&#39;0&#39;&gt;&lt;/svg&gt;" class="content_image lazy" data-actualsrc="https://pic1.zhimg.com/v2-54f656619f5520e303cd5b56e3fee8d8_b.jpg"/></figure><p>节前老师留了英语作业，手机APP上进行英语学习。</p><p>正好今天有时间，起来装好APP，跟着进行听读，然后做题。</p><p>题目都不长，边听边学，然后做选择题，很快就能完成，风格和VIPKID有些接近。</p><p>说到中秋，不得不又提起当年初一入学后，中秋节那天，小学同学和老师一起骑车去景山公园，夜里在万春亭吃月饼赏月。</p><p>大家从各自的中学汇聚到曾经的小学门前，然后骑车去景山。</p><p>每人带几块月饼，夜里在庭中大家一起分享月饼和月光。</p><p>从那之后，很多人就再没见过面了。</p><p>但是就像苏大学士说的：</p><blockquote><p>但愿人长久，千里共婵娟。</p></blockquote><hr/><p><strong><em>最近喜欢的诗</em></strong></p><blockquote><p>杜甫的天下第一律诗，对仗真是绝妙啊<br/>登高<br/>风急天高猿啸哀，渚清沙白鸟飞回。<br/>无边落木萧萧下，不尽长江滚滚来。<br/>万里悲秋常作客，百年多病独登台。<br/>艰难苦恨繁霜鬓，潦倒新停浊酒杯。</p></blockquote><hr/>' },
      recommends: [
        {id: 1, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 2, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 3, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 4, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'},
        {id: 5, title: '北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'}],
      ads: [{url: '', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'}, {url: '', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'}, {url: '', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'}, {url: '', cover: 'http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg'}]
    }
  },
  created () {
  },
  methods: {
  }
}
</script>
