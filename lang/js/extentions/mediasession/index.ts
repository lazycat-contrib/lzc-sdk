import type {MediaSessionPlugin} from './definitions';
import {MediaSessionWeb} from "./web";
declare global {
    const lzc_media_session:any;
}
if(!navigator.mediaSession){
    // @ts-ignore
    navigator.mediaSession = lzc_media_session || {};
}
const MediaSession: MediaSessionPlugin = new MediaSessionWeb()
export * from './definitions';
export {MediaSession};
