/*
 * @Author: Bin
 * @Date: 2024-10-14
 * @FilePath: /lzc-sdk/examples/lzc-react-demo/src/types.ts
 */
import { EndDeviceProxy, lzcAPIGateway } from '@lazycatcloud/sdk/index';

type PageNavigate = {
    setPageIndex: (index: number | undefined) => void;
}

type BasePageProps = {
    navigate: PageNavigate,
    gateway: lzcAPIGateway,
}

export { PageNavigate, BasePageProps }
