<template>
    <div>
        <div id="editorMenu" class="editorMenu"></div>
        <div id="editor" class="editor"></div>
    </div>
</template>

<script>
    import E from 'wangeditor'
    export default {
        name: 'editor',
        mounted () {
            var editor = new E('#editorMenu', '#editor')
            editor.create()
        }
    }
</script>

<style scoped>
    .editorMenu {
        border: 1px solid #ccc;
    }
    .editor {
        margin-top: -1px;//将多出来的一像素边框隐藏
    border: 1px solid #ccc;
        min-height: 400px;//编辑框最小高度
    height:auto;//编辑框高度超过最小高度会根据内容高度自适应
    }
</style>
