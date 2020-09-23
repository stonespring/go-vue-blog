<template>
  <a-dropdown>
    <div class="header-avatar" style="cursor: pointer">
      <a-avatar class="avatar" size="small" shape="circle" :src="pic"/>
      <span class="name">{{username}}</span>
    </div>
    <a-menu :class="['avatar-menu']" slot="overlay">
      <a-menu-item>
        <a-icon type="user" />
        <span>个人中心</span>
      </a-menu-item>
      <a-menu-item>
        <a-icon type="setting" />
        <span>设置</span>
      </a-menu-item>
      <a-menu-divider />
      <a-menu-item @click="logout">
        <a-icon style="margin-right: 8px;" type="poweroff" />
        <span>退出登录</span>
      </a-menu-item>
    </a-menu>
  </a-dropdown>
</template>

<script>
// import {mapGetters} from 'vuex'
import {logout, info} from '@/services/user'

export default {
  name: 'HeaderAvatar',
  computed: {
    // ...mapGetters('account', ['user']),
  },
  data() {
    return {
      username: "",
      pic: ""
    };
  },
  mounted() {
    this.info()
  },
  methods: {
    logout() {
      logout()
      this.$router.push('/login')
    },
    info(){
      info().then(
          res=>{
            console.log(res)
            const data = res.data.data
            this.username = data.username
            this.pic = data.pic
         }
      )
      console.log(this.username)
    }
  }
}
</script>

<style lang="less">
  .header-avatar{
    display: inline-flex;
    .avatar, .name{
      align-self: center;
    }
    .avatar{
      margin-right: 8px;
    }
    .name{
      font-weight: 500;
    }
  }
  .avatar-menu{
    width: 150px;
  }

</style>
