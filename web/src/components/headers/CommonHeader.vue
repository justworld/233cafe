<template>
  <el-header class="head1">
    <div class="logo1"></div>
    <div class="menu1"></div>
    <div class="search1">
      <el-input clearable :placeholder="searchHot"></el-input>
    </div>
    <div class="wechat1"></div>
  </el-header>
</template>

<script>
export default {
  created () {
  },
  data () {
    return {
      searchHot: 'hot'
    }
  },
  methods: {
  }
}
</script>
