import { browserHistory, hashHistory } from 'react-router'
// export default useRouterHistory(createHashHistory)()
const history = process.env.NODE_ENV === 'production' ? browserHistory : hashHistory
export default history
