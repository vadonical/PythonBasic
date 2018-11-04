import os.path
from abc import ABC

import tornado.httpserver
import tornado.ioloop
import tornado.options
import tornado.web
from tornado.options import define, options

define('port', default=8080, help='run on given port', type=int)


class IndexHandler(tornado.web.RequestHandler, ABC):
    def get(self, *args, **kwargs):
        self.render('index.html')


class PoemPageHandler(tornado.web.RequestHandler, ABC):
    def post(self):
        noun_one = self.get_argument('noun_one')
        noun_two = self.get_argument('noun_two')
        noun_thr = self.get_argument('noun_thr')
        noun_fou = self.get_argument('noun_fou')
        self.render('poem.html',
                    one=noun_one,
                    two=noun_two,
                    thr=noun_thr,
                    fou=noun_fou)


if __name__ == '__main__':
    tornado.options.parse_command_line()
    app = tornado.web.Application(
        handlers=[
            (r'/', IndexHandler),
            (r'/poem/', PoemPageHandler)
        ],
        template_path=os.path.join(os.path.dirname(__file__), 'templates')
    )
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.instance().start()
