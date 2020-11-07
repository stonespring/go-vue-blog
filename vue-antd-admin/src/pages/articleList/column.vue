<template>
    <div>
        <a-button type="primary" @click="showModal">
            添加
        </a-button>
        <a-modal
                title="添加栏目"
                :visible="visible"
                :confirm-loading="confirmLoading"
                @ok="handleOk"
                @cancel="handleCancel"
                width="40%"
        >
            <a-form :model="form" :label-col="labelCol" :wrapper-col="wrapperCol">
                <a-form-model-item label="栏目名称">
                    <a-input v-model="form.name"/>
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
            </a-form>
        </a-modal>
        <a-table bordered :data-source="dataSource" :columns="columns">

        </a-table>
    </div>
</template>
<script>

    import {column_add, column_list} from "@/services/article";

    export default {
        components: {},
        data() {
            return {
                labelCol: {span: 4},
                wrapperCol: {span: 14},
                credit_ratio: 80,
                form: {
                    name: '',
                    is_show: '',
                    top: '',
                },
                visible: false,
                confirmLoading: false,
                dataSource: [],
                count: 2,
                columns: [
                    {
                        title: 'ID',
                        dataIndex: 'id',
                        width: '30%',
                    },
                    {
                        title: '标题',
                        dataIndex: 'name',
                        width: '30%',
                        scopedSlots: {customRender: 'name'},
                    },
                    {
                        title: '是否展示',
                        dataIndex: 'is_show',
                        width: '30%',
                    },
                    {
                        title: '排序',
                        dataIndex: 'top',
                        width: '30%',
                    },
                ],
            };
        },

        mounted() {
            this.getColumnList();
        },

        methods: {
            showModal() {
                this.visible = true;
            },
            handleOk() {
                column_add(this.form).then(res => {
                    if (res.data.code == 200) {
                        this.$message.success(res.data.msg);
                        this.handleCancel() //关闭弹出框
                        this.getColumnList()
                    } else {
                        this.$message.error(res.data.msg);
                    }
                })
            },
            //获取列表
            getColumnList() {
                column_list().then(res => {
                    if (res.data.code == 200) {
                        this.dataSource = res.data.data.data
                    }
                })
            },

            handleCancel() {
                this.visible = false;
            },
            onCellChange(key, dataIndex, value) {
                const dataSource = [...this.dataSource];
                const target = dataSource.find(item => item.key === key);
                if (target) {
                    target[dataIndex] = value;
                    this.dataSource = dataSource;
                }
            },
            onDelete(key) {
                const dataSource = [...this.dataSource];
                this.dataSource = dataSource.filter(item => item.key !== key);
            },
            handleAdd() {
                const {count, dataSource} = this;
                const newData = {
                    key: count,
                    name: `Edward King ${count}`,
                    age: 32,
                    address: `London, Park Lane no. ${count}`,
                };
                this.dataSource = [...dataSource, newData];
                this.count = count + 1;
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
