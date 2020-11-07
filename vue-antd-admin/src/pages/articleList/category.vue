<template>
    <div>
        <a-button type="primary" @click="showModal">
            添加
        </a-button>
        <a-modal
                title="添加分类"
                :visible="visible"
                :confirm-loading="confirmLoading"
                @ok="categoryAdd"
                @cancel="handleCancel"
                width="40%"
        >
            <a-form-model :model="form" :label-col="labelCol" :wrapper-col="wrapperCol" @submit="onSubmit">
                <a-form-model-item label="分类名称">
                    <a-input v-model="form.category_name" />
                </a-form-model-item>
                <a-form-model-item label="栏目">
                    <a-select v-model="form.column_id" placeholder="栏目" aria-required="true">
                        <a-select-option   v-for="(item, index) in seletor" :value="item.id" :key="index">
                            {{item.name}}
                        </a-select-option>
                    </a-select>
                </a-form-model-item>

                <a-form-model-item label="是否发布">
                    <a-radio-group v-model="form.is_show">
                        <a-radio value="1">
                            是
                        </a-radio>
                        <a-radio value="2">
                            否
                        </a-radio>
                    </a-radio-group>
                </a-form-model-item>
                <a-form-model-item label="排序">
                    <a-input v-model="form.top"/>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
        <a-table bordered :data-source="dataSource" :columns="columns">
            <template slot="operation"  slot-scope="columns, record">
                <a href="javascript:;" @click="showModalTwo(record.id)" >
                    编辑
                </a>

            </template>
        </a-table>
        <a-modal
                title="编辑分类"
                :visible="visibles"
                :confirm-loading="confirmLoading"
                @ok="categoryEdit"
                @cancel="handleCancelTwo"
                width="40%"
        >
            <a-form-model :model="info" :label-col="labelCol" :wrapper-col="wrapperCol" @submit="onSubmit">
                <a-input type="hidden" v-model="info.id" />
                <a-form-model-item label="分类名称">
                    <a-input v-model="info.category_name" />
                </a-form-model-item>
                <a-form-model-item label="栏目">
                    <a-select v-model="info.column_id" placeholder="栏目" aria-required="true">
                        <a-select-option   v-for="(item, index) in seletor" :value="item.id" :key="index" >
                            {{item.name}}
                        </a-select-option>
                    </a-select>
                </a-form-model-item>

                <a-form-model-item label="是否发布">
                    <a-radio-group v-model="info.is_show">
                        <a-radio value="1">
                            是
                        </a-radio>
                        <a-radio value="2">
                            否
                        </a-radio>
                    </a-radio-group>
                </a-form-model-item>
                <a-form-model-item label="排序">
                    <a-input v-model="info.top"/>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
    </div>
</template>
<script>

    import {category_add, category_list,column_list, category_info, category_edit} from "@/services/article";

    export default {
        components: {
        },
        data() {
            return {
                labelCol: { span: 4 },
                wrapperCol: { span: 14 },
                credit_ratio: 80,
                editingKey: '',
                form: {
                    category_name: '',
                    column_id: '',
                    is_show: '',
                    top:'',
                },
                seletor: [],
                visible: false,
                visibles: false,
                confirmLoading: false,
                dataSource: [],
                count: 2,
                columns: [
                    {
                        title: 'ID',
                        dataIndex: 'id',
                    },
                    {
                        title: '分类标题',
                        dataIndex: 'category_name',
                        width: '30%',
                        scopedSlots: { customRender: 'category_name' },
                    },
                    {
                        title: '所属栏目',
                        dataIndex: 'column_id',
                    },
                    {
                        title: '是否展示',
                        dataIndex: 'is_show',
                    },
                    {
                        title: '排序',
                        dataIndex: 'top',
                    },
                    {
                        title: '操作',
                        dataIndex: 'operation',
                        scopedSlots: { customRender: 'operation' },
                    },
                ],
                info: {
                    id: 0,
                    category_name: '',
                    column_id: '',
                    is_show: '',
                    top:'',
                },
            };
        },

        mounted() {
            this.getCategoryList()
            this.getColumnList()
        },

        methods: {
            showModal() {
                this.visible = true;
            },
            showModalTwo(key) {
                this.categoryInfo(key);
                this.visibles = true;

            },
            handleCancel() {
                this.visible = false;
            },
            handleCancelTwo(){
                this.visibles = false;
            },
            categoryAdd() {
                category_add(this.form).then(res=>{
                    if (res.data.code == 200){
                        this.$message.success(res.data.msg)
                        this.handleCancel()
                        this.getCategoryList()
                    }
                })
            },
            getCategoryList(){
                category_list().then(res=>{
                    if (res.data.code == 200){
                        this.dataSource = res.data.data.data
                    }
                })
            },
            getColumnList(){
                column_list().then(res => {
                    if (res.data.code == 200) {
                        this.seletor = res.data.data.data
                    }
                })
            },

            onCellChange(key, dataIndex, value) {
                const dataSource = [...this.dataSource];
                const target = dataSource.find(item => item.key === key);
                if (target) {
                    target[dataIndex] = value;
                    this.dataSource = dataSource;
                }
            },
            categoryInfo(id){
                category_info(id).then(res=>{
                    if (res.data.code == 200){
                        this.info.id = res.data.data.data.id;
                        this.info.category_name = res.data.data.data.category_name;
                        this.info.column_id = res.data.data.data.column_id;
                        this.info.is_show = res.data.data.data.is_show;
                        this.info.top = res.data.data.data.top;
                    }else{
                        this.$message.success(res.data.msg)
                        this.handleCancelTwo()
                        this.getCategoryList()
                    }
                })
            },
            categoryEdit(){
                category_edit(this.info).then(res =>{
                    this.$message.success(res.data.msg)
                    this.handleCancelTwo()
                    this.getCategoryList()
                })
            },
            handleAdd() {
                const { count, dataSource } = this;
                const newData = {
                    key: count,
                    name: `Edward King ${count}`,
                    age: 32,
                    address: `London, Park Lane no. ${count}`,
                };
                this.dataSource = [...dataSource, newData];
                this.count = count + 1;
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
