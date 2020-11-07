import {COLUMN_ADD, COLUMN_LIST, COLUMN_EDIT,CATEGORY_INFO} from '@/services/api'
import {CATEGORY_ADD, CATEGORY_LIST, CATEGORY_EDIT} from '@/services/api'
import {ARTICLE_ADD, ARTICLE_LIST, ARTICLE_EDIT,ARTICLE_INFO} from '@/services/api'
import {request, METHOD} from '@/utils/request'

//栏目
export async function column_add(post) {
    return request(COLUMN_ADD, METHOD.POST, post)
}
export async function column_list() {
    return request(COLUMN_LIST, METHOD.GET)
}
export async function column_edit(post) {
    return request(COLUMN_EDIT, METHOD.POST, post)
}

//todo 分类
export async function category_add(post) {
    return request(CATEGORY_ADD, METHOD.POST, post)
}
//列表
export async function category_list() {
    return request(CATEGORY_LIST, METHOD.GET)
}
//编辑提交
export async function category_edit(post) {
    return request(CATEGORY_EDIT, METHOD.POST, post)
}
export async function category_info(id) {
    return request(CATEGORY_INFO, METHOD.GET, {
        id: id
    })
}


//todo  文章
export async function article_add(post) {
    return request(ARTICLE_ADD, METHOD.POST, post)
}
//列表
export async function article_list() {
    return request(ARTICLE_LIST, METHOD.GET)
}
//编辑提交
export async function article_edit(post) {
    return request(ARTICLE_EDIT, METHOD.POST, post)
}
export async function article_info(id) {
    return request(ARTICLE_INFO, METHOD.GET, {
        id: id
    })
}
