import { grpc } from "@improbable-eng/grpc-web"
import { BrowserHeaders } from "browser-headers"
import { share } from "rxjs/operators"
import { Observable } from "rxjs"
import _m0 from "protobufjs/minimal"

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
    upStreamRetryCodes?: number[]
  }

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory
      streamingTransport?: grpc.TransportFactory
      debug?: boolean
      metadata?: grpc.Metadata
      upStreamRetryCodes?: number[]
    },
  ) {
    this.host = host
    this.options = options
  }

  private addExtHeaders(meta: BrowserHeaders) {
    // meta可能不存在
    meta?.set("x-request-time", new Date().getTime() + "")
  }

  unary<T extends UnaryMethodDefinitionish>(methodDesc: T, _request: any, metadata: grpc.Metadata | undefined, abortSignal?: AbortSignal): Promise<any> {
    const request = { ..._request, ...methodDesc.requestType }
    const maybeCombinedMetadata =
      metadata && this.options.metadata ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap }) : metadata ?? this.options.metadata
    this.addExtHeaders(maybeCombinedMetadata)
    return new Promise((resolve, reject) => {
      let that = this
      const client = grpc.unary(methodDesc, {
        request,
        host: this.host,
        metadata: maybeCombinedMetadata ?? {},
        ...(this.options.transport !== undefined ? { transport: this.options.transport } : {}),
        debug: this.options.debug,
        onEnd: function (response) {
          if (that.options.debug) console.table(response)
          if (response.status === grpc.Code.OK) {
            resolve(response.message!.toObject())
          } else if (response.status == grpc.Code.Unauthenticated) {
            console.warn("sdk 401 auto handle")
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
            const err = new GrpcWebError(response.statusMessage, response.status, response.trailers)
            reject(err)
          }
        },
      })
      if (abortSignal) {
        abortSignal.addEventListener("abort", () => {
          client.close()
          reject(abortSignal.reason)
        })
      }
    })
  }

  invoke<T extends UnaryMethodDefinitionish>(methodDesc: T, _request: any, metadata: grpc.Metadata | undefined, abortSignal?: AbortSignal): Observable<any> {
    const upStreamCodes = this.options.upStreamRetryCodes ?? []
    const DEFAULT_TIMEOUT_TIME: number = 3_000
    const request = { ..._request, ...methodDesc.requestType }
    const transport = this.options.streamingTransport ?? this.options.transport
    const maybeCombinedMetadata =
      metadata && this.options.metadata ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap }) : metadata ?? this.options.metadata
    this.addExtHeaders(maybeCombinedMetadata)
    return new Observable(observer => {
      const upStream = () => {
        const client = grpc.invoke(methodDesc, {
          host: this.host,
          request,
          ...(transport !== undefined ? { transport } : {}),
          metadata: maybeCombinedMetadata ?? {},
          debug: this.options.debug ?? false,
          onMessage: next => observer.next(next),
          onEnd: (code: grpc.Code, message: string, trailers: grpc.Metadata) => {
            if (code === 0) {
              observer.complete()
            } else if (upStreamCodes.includes(code)) {
              setTimeout(upStream, DEFAULT_TIMEOUT_TIME)
            } else {
              const err = new Error(message) as any
              err.code = code
              err.metadata = trailers
              observer.error(err)
            }
          },
        })

        if (abortSignal) {
          const abort = () => {
            observer.error(abortSignal.reason)
            client.close()
          }
          abortSignal.addEventListener("abort", abort)
          observer.add(() => {
            if (abortSignal.aborted) {
              return
            }

            abortSignal.removeEventListener("abort", abort)
            client.close()
          })
        } else {
          observer.add(() => client.close())
        }
      }
      upStream()
    }).pipe(share())
  }
}

export class GrpcWebError extends globalThis.Error {
  constructor(
    message: string,
    public code: grpc.Code,
    public metadata: grpc.Metadata,
  ) {
    super(message)
  }
}
