<template>
    <div class="col-md-9 left-col">
        <div class="panel panel-default padding-md">
            <div class="panel-body">
                <h2><i class="fa fa-lock"></i> 修改密码</h2>
                <hr>
                <div class="form-horizontal" data-validator-form>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">密 码</label>
                        <div class="col-sm-6">
                            <input v-model.trim="password" id="password" v-validator.required="{ regex: /^\w{6,16}$/, error: '密码要求 6 ~ 16 个单词字符' }" type="password" class="form-control" placeholder="请填写密码">
                        </div>
                    </div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">确认密码</label>
                        <div class="col-sm-6">
                            <input v-model.trim="cpassword" v-validator.required="{ title: '确认密码', target: '#password' }" type="password" class="form-control">
                        </div>
                    </div>

                    <div class="form-group">
                        <div class="col-sm-offset-2 col-sm-6">
                            <button type="submit" class="btn btn-primary" @click="updatePassword">应用修改</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'EditPassword',
        data() {
            return {
                password: '', // 密码
                cpassword: '' // 确认密码
            }
        },
        // 在实例创建完成后，初始化密码的值
        created() {
            // 获取仓库的用户信息
            const user = this.$store.state.user

            if (user && typeof user === 'object') {
                // 将仓库的用户密码赋值给当前密码
                this.password = user.password
            }
        },
        methods: {
            // 更新密码
            updatePassword(e) {
                this.$nextTick(() => {
                    // 表单验证通过时
                    if (e.target.canSubmit) {
                        // 依然分发一个 updateUser 的事件，这里只需传入密码信息
                        this.$store.dispatch('updateUser', { password: this.cpassword })
                        this.$message.show('修改成功')
                    }
                })
            }
        }
    }
</script>

<style scoped>

</style>
