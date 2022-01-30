// @flow

// https://stackoverflow.com/questions/3665115/how-to-create-a-file-in-memory-for-user-to-download-but-not-through-server
export default function downloadText(content: string, filename: string) {
  const a = document.createElement("a");
  const blob = new Blob([content], { type: "text/plain" });
  const url = URL.createObjectURL(blob);
  a.setAttribute("href", url);
  a.setAttribute("download", filename);
  a.click();
}
