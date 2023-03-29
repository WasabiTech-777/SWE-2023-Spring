import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpHandler, HttpRequest } from '@angular/common/http';

@Injectable()
export class InterceptorService implements HttpInterceptor {

  intercept(req: HttpRequest<any>, next: HttpHandler) {

    const corsReq = req.clone({
      headers: req.headers.set('Access-Control-Allow-Origin', '*')
      .set("Access-Control-Allow-Credentials", 'true')
    });

    return next.handle(corsReq);
  }
}
