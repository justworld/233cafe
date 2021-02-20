<template>
  <el-container direction="vertical">
    <router-view name="header"></router-view>
    <router-view name="content"></router-view>
    <router-view name="footer"></router-view>
    <el-backtop class="backtop1" target="#app"></el-backtop>
  </el-container>
</template>

<script>
</script>

<style lang="stylus" rel="stylesheet/stylus">
  @import "../common/stylus/base.styl";
</style>
