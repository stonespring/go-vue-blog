<template>
    <div>
        <a-button type="primary" @click="showModal">
            添加
        </a-button>
        <a-modal
                title="添加文章"
                :visible="visible"
                :confirm-loading="confirmLoading"
                @ok="handleOk"
                @cancel="handleCancel"
                width="40%"
        >
            <a-form-model :model="form" :label-col="labelCol" :wrapper-col="wrapperCol">
                <a-form-model-item label="文章名称">
                    <a-input v-model="form.name" />
                </a-form-model-item>
                <a-form-model-item label="分类">
                    <a-select v-model="form.region" placeholder="类别">
                        <a-select-option value="shanghai">
                            Zone one
                        </a-select-option>
                        <a-select-option value="beijing">
                            Zone two
                        </a-select-option>
                    </a-select>
                </a-form-model-item>
                <a-form-model-item label="创建时间">
                    <a-date-picker
                            v-model="form.date1"
                            show-time
                            type="date"
                            placeholder="时间"
                            style="width: 100%;"
                    />
                </a-form-model-item>

                <a-form-model-item label="是否发布">
                    <a-radio-group v-model="form.resource">
                        <a-radio value="1">
                            是
                        </a-radio>
                        <a-radio value="2">
                            否
                        </a-radio>
                    </a-radio-group>
                </a-form-model-item>
                <a-form-model-item label="内容">
                    <a-input v-model="form.desc" type="textarea" />
                </a-form-model-item>
<!--                <a-form-model-item :wrapper-col="{ span: 14, offset: 4 }">-->
<!--                    <a-button type="primary" @click="onSubmit">-->
<!--                        创建-->
<!--                    </a-button>-->
<!--                    <a-button style="margin-left: 10px;">-->
<!--                        取消-->
<!--                    </a-button>-->
<!--                </a-form-model-item>-->
            </a-form-model>
        </a-modal>
        <a-table bordered :data-source="dataSource" :columns="columns">
            <template slot="name" slot-scope="text, record">
                <editable-cell :text="text" @change="onCellChange(record.key, 'name', $event)" />
            </template>
            <template slot="operation" slot-scope="text, record">
                <a-popconfirm
                        v-if="dataSource.length"
                        title="Sure to delete?"
                        @confirm="() => onDelete(record.key)"
                >
                    <a href="javascript:;">删除</a>
                </a-popconfirm>
            </template>
        </a-table>
    </div>
</template>
<script>

    export default {
        components: {
        },
        data() {
            return {
                labelCol: { span: 4 },
                wrapperCol: { span: 14 },
                credit_ratio: 80,
                form: {
                    name: '',
                    region: undefined,
                    date1: undefined,
                    delivery: false,
                    type: [],
                    resource: '',
                    desc: '',
                },
                visible: false,
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
                        title: '标题',
                        dataIndex: 'name',
                        width: '30%',
                        scopedSlots: { customRender: 'name' },
                    },
                    {
                        title: '创建时间',
                        dataIndex: 'create_time',
                    },
                    {
                        title: '操作',
                        dataIndex: 'operation',
                        scopedSlots: { customRender: 'operation' },
                    },
                ],
            };
        },
        methods: {
            showModal() {
                this.visible = true;
            },
            handleOk() {
                this.confirmLoading = true;
                setTimeout(() => {
                    this.visible = false;
                    this.confirmLoading = false;
                }, 2000);
            },
            handleCancel() {
                console.log('Clicked cancel button');
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
