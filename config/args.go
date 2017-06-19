package config

type Args []string

func parseArgs(args *[]string) *Args {
	if args == nil || len(*args) == 0 {
		return &Args{}
	}

	result := Args(*args)
	return &result
}

func (a *Args) prependArgs(newArgs ...string) *Args {
	count := len(newArgs)
	result := make(Args, count+len(*a))
	for k, v := range newArgs {
		result[k] = v
	}
	for k, v := range *a {
		result[count+k] = v
	}
	return &result
}
