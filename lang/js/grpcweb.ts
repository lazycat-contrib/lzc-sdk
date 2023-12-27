import { of, Observable } from "rxjs"
import Long from "long"
import { grpc } from "@improbable-eng/grpc-web"
import { BrowserHeaders } from "browser-headers"
import { share } from "rxjs/operators"

interface UnaryMethodDefinitionishR extends grpc.UnaryMethodDefinition<any, any> {
  requestStream: any
  responseStream: any
}

type UnaryMethodDefinitionish = UnaryMethodDefinitionishR

export class GrpcWebImpl {
  private host: string
  private options: {
    transport?: grpc.TransportFactory
    streamingTransport?: grpc.TransportFactory
    debug?: boolean
    metadata?: grpc.Metadata
  }

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory
      streamingTransport?: grpc.TransportFactory
      debug?: boolean
      metadata?: grpc.Metadata
    },
  ) {
    this.host = host
    this.options = options
  }

  unary<T extends UnaryMethodDefinitionish>(methodDesc: T, _request: any, metadata: grpc.Metadata | undefined): Promise<any> {
    const request = {..._request, ...methodDesc.requestType}
    const maybeCombinedMetadata =
        metadata && this.options.metadata
            ? new BrowserHeaders({
              ...this.options?.metadata.headersMap,
              ...metadata?.headersMap,
            })
            : metadata || this.options.metadata
    this.addExtHeaders(maybeCombinedMetadata);
    return new Promise((resolve, reject) => {
      let that = this
      grpc.unary(methodDesc, {
        request,
        host: this.host,
        metadata: maybeCombinedMetadata,
        transport: this.options.transport,
        debug: this.options.debug,
        onEnd: function (response) {
          if (that.options.debug) console.table(response)
          if (response.status === grpc.Code.OK) {
            resolve(response.message)
            // auto handle 401 error
          } else if (response.status === grpc.Code.Unauthenticated) {
            console.log("sdk 401 auto handle")
            const domainList = window.location.host.split(".")
            let boxdomain = ""
            // 忽略boxdomain下的子域名
            if (domainList.length == 3) {
              boxdomain = window.location.host
            } else {
              boxdomain = domainList.slice(domainList.length - 3).join(".")
            }
            window.location.replace(`${window.location.protocol}//${boxdomain}/sys/login?redirect=${encodeURIComponent(window.location.href)}`)
            const err = new Error(response.statusMessage) as any
            err.code = response.status
            err.metadata = response.trailers
            reject(err)
          } else {
            const err = new Error(response.statusMessage) as any
            err.code = response.status
            err.metadata = response.trailers
            reject(err)
          }
        },
      })
    })
  }

  private addExtHeaders(meta: BrowserHeaders) {
    // meta可能不存在
    meta?.set("x-request-time", new Date().getTime() + "")
  }

  invoke<T extends UnaryMethodDefinitionish>(methodDesc: T, _request: any, metadata: grpc.Metadata | undefined): Observable<any> {
    const request = { ..._request, ...methodDesc.requestType }
    const maybeCombinedMetadata =
      metadata && this.options.metadata
        ? new BrowserHeaders({
            ...this.options?.metadata.headersMap,
            ...metadata?.headersMap,
          })
        : metadata || this.options.metadata
    this.addExtHeaders(maybeCombinedMetadata)
    return new Observable(observer => {
      const upStream = () => {
        const client = grpc.invoke(methodDesc, {
          host: this.host,
          request,
          transport: this.options.streamingTransport || this.options.transport,
          metadata: maybeCombinedMetadata,
          debug: this.options.debug,
          onMessage: next => observer.next(next),
          onEnd: (code: grpc.Code, message: string) => {
            if (code === 0) {
              observer.complete()
            } else {
              observer.error(new Error(`Error ${code} ${message}`))
            }
          },
        })
        observer.add(() => client.close())
      }
      upStream()
    }).pipe(share())
  }
}
