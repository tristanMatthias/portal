import { EventsEmit, EventsOn } from '../../wailsjs/runtime'

export function downloadModel(model: string, cb: (data: any) => void) {
  EventsEmit("download", model);
  EventsOn("download-progress", cb);
}
