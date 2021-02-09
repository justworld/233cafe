<template>
  <el-header class="head1">
    <div class="head2">
      <div class="logo1"></div>
      <div class="menu1"></div>
      <div class="search1">
        <el-input clearable :placeholder="searchHot"></el-input>
      </div>
      <div class="wechat1"></div>
    </div>
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
