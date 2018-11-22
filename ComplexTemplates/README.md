# ComplexTemplates

### Tips

参看下列代码： 

```text
handlers=[
            (r'/', IndexHandler),
            (r'/poem', MungePageHandler)
        ],
template_path=os.path.join(os.path.dirname(__file__), 'templates'),
static_path=os.path.join(os.path.dirname(__file__), 'static'),
debug=True
```

- 每个路由对应相应的类。
- template_path 表示设置模板文件的路径。其中放置网页模板，支持模板继承。
- static_path 表示静态属性的路径，其中也应该放一些样式、环境包等内容。
- 在模板中的 HTML 文档可以继承和替换，实现了代码的重用性。
- 导入静态文件时应该是下列操作：

```text
<link rel="stylesheet" href="{{ static_url('style.css') }}">
```

- 关键字是 static_url。