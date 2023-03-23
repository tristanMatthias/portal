export function downloadModel(model: string, cb: (data: any) => void) {
  window.runtime.EventsEmit("download", model);
  window.runtime.EventsOn("download-progress", cb);
}
