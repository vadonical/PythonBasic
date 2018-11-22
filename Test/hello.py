import tornado.web
import tornado.ioloop
import tornado.options
import tornado.httpserver

from tornado.options import define, options

define('port', default=8081, help='Run on the port 8080', type=int)


class IndexHandler(tornado.web.RequestHandler):
    def get(self):
        greet = self.get_argument('greet', 'Hello')
        self.write(greet + ', My best friend!')


if __name__ == '__main__':
    tornado.options.parse_command_line()
    # handler后面应该是个元组
    app = tornado.web.Application(handlers=[(r"/", IndexHandler)])
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.instance().start()
