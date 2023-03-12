const baseUrl="http://127.0.0.1:8080"
const url = `${baseUrl}/api`;

export const daliurenAPI :{url: string, header: object, method: "POST"}= {
  url: `${url}/daliuren`,
  header: { "Content-Type": "application/json" },
  method: "POST",
};
