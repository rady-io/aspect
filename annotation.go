package aspect

const (
	/*
		BeforeAnn is a annotation to invoke a method/func before invoking this method/func.
		AfterAnn is a annotation to invoke a method/func after invoking this method/func.

		// @Before srv.log.Debugf("invoke method {name} in {file}:{position}, deliver %s", str)
		// @After srv.log.Debugf("invoked {func} successfully")
		func (srv *Server) exampleMethod(str string) {
			...
		}

		if the results of this method/func contain error:
			- before method/func can have only one result whose type is error,
			and the whole invoking will return if the that is not nil.
			- after method/func will never be invoked if this error is not nil.
	*/
	BeforeAnn = "@Before"
	AfterAnn  = "@After"

	/*
		ProxyAnn is a annotation to filter method/func, like decorator in python.

		type Data struct {
			...
		}

		func CacheProxy(getter func(id int, name string) (*Data, error)) func(id int, name string) (*Data, error) {
			return func(id int, name string) (data *Data, err error) {
				key := fmt.Sprintf("%d-%s", id, name)
				if data = cache.Get(key); data != nil {
					if data, err = getter(id, name); err == nil {
						cache.Set(key, data)
					}
				}
				return
			}
		}

		// @Proxy(CacheProxy)
		func getData(id int, name string) (*Data, error) {
			...
		}
	*/
	ProxyAnn = "@Proxy"
)
