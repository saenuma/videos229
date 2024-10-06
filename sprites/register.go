package sprites

func RegisterAll(tmplStore *map[string]string) {
	(*tmplStore)["sprites/1"] = Method1Conf
	// (*tmplStore)["sprites/2"] = Method2Conf
	(*tmplStore)["sprites/3"] = Method3Conf
	(*tmplStore)["sprites/4"] = Method4Conf
	(*tmplStore)["sprites/5"] = Method5Conf
}
