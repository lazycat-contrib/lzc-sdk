import type {
    ActionHandler,
    ActionHandlerOptions,
    MediaSessionPlugin,
    MetadataOptions,
    PlaybackStateOptions,
    PositionStateOptions
} from './definitions';

const functionMap: Map<string, ActionHandler> = new Map();
window.addEventListener('lzc_media_session_event', (e: any) => {
    const execFun = functionMap.get(e.detail.eventType)
    if (execFun) {
        execFun(e.detail.data);
    }
});

export class MediaSessionWeb implements MediaSessionPlugin {
    async setMetadata(options: MetadataOptions): Promise<void> {
        if (lzc_media_session) {
            // @ts-ignore
            navigator.mediaSession.setMetadata(JSON.stringify(options));
        } else if ('mediaSession' in navigator) {
            navigator.mediaSession.metadata = new MediaMetadata(options);
        } else {
            // @ts-ignore
            throw this.unavailable('Media Session API not available in this browser.');
        }
    }

    async setPlaybackState(options: PlaybackStateOptions): Promise<void> {
        if (lzc_media_session) {
            // @ts-ignore
            navigator.mediaSession.setPlaybackState(options.playbackState);
        } else if ('mediaSession' in navigator) {
            navigator.mediaSession.playbackState = options.playbackState;
        } else {
            // @ts-ignore
            throw this.unavailable('Media Session API not available in this browser.');
        }
    };

    async setActionHandler(options: ActionHandlerOptions, handler: ActionHandler | null): Promise<void> {
        if (lzc_media_session) {
            // 单独处理
            // @ts-ignore
            functionMap.set(options.action, handler);
            // @ts-ignore
            navigator.mediaSession.setActionHandler(options.action)
        } else if ('mediaSession' in navigator) {
            navigator.mediaSession.setActionHandler(options.action, handler);
        } else {
            // @ts-ignore
            throw this.unavailable('Media Session API not available in this browser.');
        }
    };

    async setPositionState(options: PositionStateOptions): Promise<void> {
        if (lzc_media_session) {
            // @ts-ignore
            navigator.mediaSession.setPositionState(JSON.stringify(options));
        }
        if ('mediaSession' in navigator) {
            navigator.mediaSession.setPositionState(options);
        } else {
            // @ts-ignore
            throw this.unavailable('Media Session API not available in this browser.');
        }
    };
}
