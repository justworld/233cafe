<template>
  <el-header class="head1">
    <div class="logo1" @click="goHome"></div>
    <div class="menu1">
      <el-link :underline="false" href="/j" class="link1 crumb1" :class="{'menu-active1': 'j'==active}">经济</el-link>
      <el-link :underline="false" href="/k" class="link1 crumb1" :class="{'menu-active1': 'k'==active}">科技</el-link>
      <el-link :underline="false" href="/s" class="link1 crumb1" :class="{'menu-active1': 's'==active}">生活</el-link>
      <el-link :underline="false" href="/t" class="link1 crumb1" :class="{'menu-active1': 't'==active}">体育</el-link>
      <el-link :underline="false" href="/y" class="link1 crumb1" :class="{'menu-active1': 'y'==active}">娱乐</el-link>
    </div>
    <div class="search1">
      <el-input clearable :placeholder="searchHot"><i slot="suffix" class="el-input__icon el-icon-search"></i></el-input>
    </div>
    <div class="wechat1">
      <el-dropdown class="dropdown1" placement="top-end">
        <span>关注微信 <i class="el-icon-caret-bottom"></i></span>
        <el-dropdown-menu slot="dropdown" class="dropdown1-menu1">
          <el-dropdown-item>
            jlkjlkjlkj
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </el-header>
</template>

<script>
export default {
  created () {
    let $this = this
    switch ($this.$route.meta.menuIndex) {
      case 'j':
        $this.active = 'j'
        break
      case 'k':
        $this.active = 'k'
        break
      case 's':
        $this.active = 's'
        break
      case 't':
        $this.active = 't'
        break
      case 'y':
        $this.active = 'y'
        break
      default:
        $this.active = ''
        break
    }
  },
  data () {
    return {
      searchHot: 'hot'
    }
  },
  methods: {
    goHome () {
      location.href = '/'
    }
  }
}
</script>
