import os
import tornado.httpserver
import tornado.ioloop
import tornado.web
import tornado.options
from tornado.options import define, options

define('port', default=8082, help='Run this server', type=int)


class IndexHandler(tornado.web.RequestHandler):
    def get(self, *args, **kwargs):
        self.render('index.html')


class PoemMaker(tornado.web.RequestHandler):
    def post(self, *args, **kwargs):
        str1 = self.get_argument('str1')
        str2 = self.get_argument('str2')
        str3 = self.get_argument('str3')
        str4 = self.get_argument('str4')
        self.render('poem.html', one=str1, two=str2, thr=str3, fou=str4)


if __name__ == '__main__':
    tornado.options.parse_command_line()
    app = tornado.web.Application(
        handlers=[(r'/', IndexHandler),
                  (r'/poem', PoemMaker)],
        template_path=os.path.join(os.path.dirname(__file__), 'templates')
    )
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.instance().start()
