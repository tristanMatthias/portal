import { ActionDownload } from '../../wailsjs/go/model/model';
import { EventsOn } from '../../wailsjs/runtime';

export function downloadModel(model: string, cb: (data: any) => void) {
  ActionDownload(model);
  EventsOn("download-progress", cb);
}
