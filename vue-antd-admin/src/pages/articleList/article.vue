<template>
    <div>
        <a-button type="primary" @click="showModal">
            添加
        </a-button>
        <a-modal
                title="添加文章"
                :visible="visible"
                :confirm-loading="confirmLoading"
                @ok="articleAdd"
                @cancel="handleCancel"
                width="70%"
        >
            <a-form-model :model="form" :label-col="labelCol" :wrapper-col="wrapperCol" @submit="onSubmit">
                <a-form-model-item label="文章名称">
                    <a-input v-model="form.title" aria-required="true"/>
                </a-form-model-item>
                <a-form-model-item label="分类">
                    <a-select v-model="form.category_id" placeholder="类别" aria-required="true">
                        <a-select-option v-for="(item, index) in category_list" :value="item.id" :key="index">
                            {{item.category_name}}
                        </a-select-option>

                    </a-select>
                </a-form-model-item>
                <a-form-model-item label="创建时间">
                    <a-date-picker
                            v-model="form.create_date"
                            show-time
                            type="date"
                            placeholder="时间"
                            style="width: 100%;"
                            aria-required="true"
                    />
                </a-form-model-item>


                <a-form-model-item label="文章描述">
                    <a-input v-model="form.desc" aria-required="true"/>
                </a-form-model-item>
                <a-form-model-item label="内容"  >
                    <div>
                        <quill-editor ref="myTextEditor" v-model="form.content" :options="editorOption"  width="100%" aria-required="true"></quill-editor>
                    </div>
                </a-form-model-item>
                <a-form-model-item label="是否发布">
                    <a-radio-group v-model="form.is_show" aria-required="true">
                        <a-radio value="1">
                            是
                        </a-radio>
                        <a-radio value="2">
                            否
                        </a-radio>
                    </a-radio-group>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
        <a-table bordered :data-source="dataSource" :columns="columns">
            <template slot="operation" slot-scope="text, record">
                    <a href="javascript:;" @click="showModalTwo(record.id)" >
                        编辑
                    </a>

                <a-popconfirm
                        v-if="dataSource.length"
                        title="确定删除?"
                        @confirm="() => onDelete(record.key)"
                >
                    <a href="javascript:;">删除</a>
                </a-popconfirm>
            </template>
        </a-table>
        <a-modal
                title="编辑文章"
                :visible="visibleTwo"
                :confirm-loading="confirmLoading"
                @ok="articleEdit"
                @cancel="handleCancel"
                width="70%"
        >
            <a-form-model :model="form" :label-col="labelCol" :wrapper-col="wrapperCol" @submit="onSubmit">
                <a-form-model-item label="文章名称">
                    <a-input v-model="form.title" aria-required="true"/>
                </a-form-model-item>
                <a-form-model-item label="分类">
                    <a-select v-model="form.category_id" placeholder="类别" aria-required="true">
                        <a-select-option v-for="(item, index) in category_list" :value="item.id" :key="index">
                            {{item.category_name}}
                        </a-select-option>

                    </a-select>
                </a-form-model-item>
                <a-form-model-item label="创建时间">
                    <a-date-picker
                            v-model="form.create_date"
                            show-time
                            type="date"
                            placeholder="时间"
                            style="width: 100%;"
                    />
                </a-form-model-item>


                <a-form-model-item label="文章描述">
                    <a-input v-model="form.desc" aria-required="true"/>
                </a-form-model-item>
                <a-form-model-item label="内容"  >
                    <div>
                        <quill-editor ref="myTextEditor" v-model="form.content" :options="editorOption"  width="100%" aria-required="true"></quill-editor>
                    </div>
                </a-form-model-item>
                <a-form-model-item label="是否发布">
                    <a-radio-group v-model="form.is_show" aria-required="true">
                        <a-radio value="1">
                            是
                        </a-radio>
                        <a-radio value="2">
                            否
                        </a-radio>
                    </a-radio-group>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
    </div>
</template>

<script>
    import 'quill/dist/quill.core.css'
    import 'quill/dist/quill.snow.css'
    import 'quill/dist/quill.bubble.css'

    import { quillEditor } from 'vue-quill-editor'
    import {article_add,article_list, category_list,article_edit,article_info} from "@/services/article";
    export default {
        components: {
            quillEditor
        },
        data() {
            return {
                //文本编辑器
                editorOption: {
                    placeholder: '书写文章内容'
                },
                //其他
                labelCol: { span: 4 },
                wrapperCol: { span: 14 },
                credit_ratio: 80,
                category_list: [],
                form: {
                    title: '',
                    category_id: undefined,
                    create_date: undefined,
                    is_show: 1,
                    desc: '',
                    content: '',
                },
                visible: false,
                visibleTwo: false,
                confirmLoading: false,
                dataSource: [
                    {
                        key: '0',
                        name: 'Edward King 0',
                        age: '32',
                    },
                    {
                        key: '1',
                        name: 'Edward King 1',
                        create_time: '32',
                    },
                ],
                count: 2,
                columns: [
                    {
                        title: 'ID',
                        dataIndex: 'id',
                    },
                    {
                        title: '标题',
                        dataIndex: 'title',
                        width: '20%',
                        scopedSlots: { customRender: 'title' },
                    },
                    {
                        title: '所属分类',
                        dataIndex: 'category_id',
                    },
                    {
                        title: '文章描述',
                        dataIndex: 'desc',
                        width: '20%',
                        scopedSlots: { customRender: 'desc' },
                    },
                    {
                        title: '是否发布',
                        dataIndex: 'is_show',
                    },
                    {
                        title: '创建时间',
                        dataIndex: 'create_date',
                    },
                    {
                        title: '操作',
                        dataIndex: 'operation',
                        scopedSlots: { customRender: 'operation' },
                    },
                ],
            };
        },
        mounted() {
            this.getArticleList();
        },
        methods: {
            //获取分类数据
            getCategoryList(){
                category_list().then(res=>{
                    if (res.data.code == 200){
                        this.category_list = res.data.data.data
                    }
                })
            },

            getArticleList(){
              article_list().then(res=>{
                  if (res.data.code == 200){
                      this.dataSource = res.data.data;
                  }
              })
            },

            showModal() {
                this.visible = true;
                this.form.content= ''; //默认文章内容为空
                this.getCategoryList();
            },
            showModalTwo(key){
                this.articleInfo(key);
                this.visibleTwo = true; //模态框弹出
                this.getCategoryList();
            },
            articleAdd() {
                article_add(this.form).then(res => {
                    if (res.data.code == 200){
                        this.$message.success(res.data.msg);
                        this.handleCancel()

                    }
                })
            },
            articleEdit() {
                article_edit(this.form).then(res => {
                    if (res.data.code == 200){
                        this.$message.success(res.data.msg);
                        this.handleCancel()

                    }
                })
            },
            articleInfo(id){
                article_info(id).then(res =>{
                    if (res.data.code == 200){
                        this.form.id = res.data.data.id;
                        this.form.category_id = res.data.data.category_id;
                        this.form.title = res.data.data.title;
                        this.form.create_date = res.data.data.create_date;
                        this.form.is_show = res.data.data.is_show;
                        this.form.content = res.data.data.content;
                        this.form.desc = res.data.data.desc;
                    }
                })
            },
            handleCancel() {
                this.visible = false;
                this.visibleTwo = false;
            },
            //删除操作暂时不弄了..自己玩
            onDelete(key) {
                console.log(key);
                // const dataSource = [...this.dataSource];
                // this.dataSource = dataSource.filter(item => item.key !== key);
            },

            onSubmit (e) {
                e.preventDefault()
            },
        },
    };
</script>
<style>
    .editable-cell {
        position: relative;
    }

    .editable-cell-input-wrapper,
    .editable-cell-text-wrapper {
        padding-right: 24px;
    }

    .editable-cell-text-wrapper {
        padding: 5px 24px 5px 5px;
    }

    .editable-cell-icon,
    .editable-cell-icon-check {
        position: absolute;
        right: 0;
        width: 20px;
        cursor: pointer;
    }

    .editable-cell-icon {
        line-height: 18px;
        display: none;
    }

    .editable-cell-icon-check {
        line-height: 28px;
    }

    .editable-cell:hover .editable-cell-icon {
        display: inline-block;
    }

    .editable-cell-icon:hover,
    .editable-cell-icon-check:hover {
        color: #108ee9;
    }

    .editable-add-btn {
        margin-bottom: 8px;
    }
</style>
