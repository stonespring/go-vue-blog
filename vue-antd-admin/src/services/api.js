//跨域代理前缀
const API_PROXY_PREFIX='http://localhost:8080/admin'
const BASE_URL = process.env.NODE_ENV === 'production' ? process.env.VUE_APP_API_BASE_URL : API_PROXY_PREFIX
// const BASE_URL = process.env.VUE_APP_API_BASE_URL
module.exports = {
  LOGIN: `${BASE_URL}/login`,
  ROUTES: `${BASE_URL}/v1/routes`,
  INFO: `${BASE_URL}/v1/info`,

  // 栏目添加
  COLUMN_ADD: `${BASE_URL}/v1/column_add`,
  //栏目列表
  COLUMN_LIST: `${BASE_URL}/v1/column_list`,
  //栏目编辑展示
  COLUMN_EDIT: `${BASE_URL}/v1/column_edit`,

  // 栏目添加
  CATEGORY_ADD: `${BASE_URL}/v1/category_add`,
  //栏目列表
  CATEGORY_LIST: `${BASE_URL}/v1/category_list`,
  //栏目编辑展示
  CATEGORY_INFO: `${BASE_URL}/v1/category_info`,
  //栏目编辑展示
  CATEGORY_EDIT: `${BASE_URL}/v1/category_edit`,


  // 文章添加
  ARTICLE_ADD: `${BASE_URL}/v1/article_add`,
  //文章列表
  ARTICLE_LIST: `${BASE_URL}/v1/article_list`,
  //文章编辑展示
  ARTICLE_INFO: `${BASE_URL}/v1/article_info`,
  //文章编辑展示
  ARTICLE_EDIT: `${BASE_URL}/v1/article_edit`,

}
