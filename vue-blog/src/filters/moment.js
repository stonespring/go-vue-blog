// 引入 moment.js 的默认值
import moment from 'moment'

// 设置语言为中文
moment.locale('zh-cn')

export default function(value, ...rest) {
    // 获取第一个日期参数
    const date = value

    // 验证日期是否合法
    if (moment(date).isValid()) {
        // 获取第二个参数
        const key = rest.shift()

        if (typeof key === 'string') {
            switch (key) {
                case 'from':
                    // 格式化日期为多久之前
                    value = moment(date).from()

                    // 获取第三个参数
                    const otherOpts = rest.shift()

                    if (otherOpts && typeof otherOpts === 'object') {
                        // 如果参数对象有 startOf 属性，就使用 startOf 的值作为开始时间
                        value = moment(date).startOf(otherOpts.startOf).from()
                    }

                    break
                default:
                    // 直接使用第二个参数进行格式化
                    value = moment(date).format(key)
            }
        }
    }

    // 返回处理后的日期
    return value
}
