import {withRouter} from "react-router-dom";
import "ces-theme/dist/css/ces.css";
import {QueryClient, QueryClientProvider} from "react-query";

const queryClient = new QueryClient()

function App() {
    return (
        <QueryClientProvider client={queryClient}>
            <div style={{"paddingTop": "60px", "paddingLeft": "300px", "width": "80%"}}>
                <p>MyApp</p>
            </div>
        </QueryClientProvider>
    );
}

export default withRouter(App);
