module.exports = {
    title: 'BBS-GO',
    description: '基于Go语言的开源BBS系统',
    themeConfig: {
        search: true,
        nav: [
            {text: 'Home', link: '/'},
            {text: 'Guide', link: '/guide/'},
            {text: '码农俱乐部', link: 'https://mlog.club'},
        ],

        // 以下为可选的编辑链接选项
        // 假如你的文档仓库和项目本身不在一个仓库：
        repo: 'mlogclub/bbs-go',
        // 假如文档不是放在仓库的根目录下：
        docsDir: 'docs/docs',
        // 假如文档放在一个特定的分支下：
        docsBranch: 'master',
        // 默认是 false, 设置为 true 来启用
        editLinks: false,
        // 默认为 "Edit this page"
        editLinkText: '帮助我们改善此页面！'
    }
}