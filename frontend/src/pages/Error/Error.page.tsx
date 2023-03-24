import { useRouteError } from "react-router-dom";
import Page from "../../components/Page/Page";

export default function PageError() {
  const error = useRouteError() as any;
  console.error(error);

  return (
    <Page id="error" title="Oops!">
      <p>Sorry, an unexpected error has occurred.</p>
      <p>
        <i>{error.statusText || error.message}</i>
      </p>
    </Page>
  );
}
