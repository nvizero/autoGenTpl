package control

var (
	project_name string = "isb"
	port_serial  int    = 55
	No           int    = 33
	status       chan string
	params       = []interface{}{}
	cmd_sh       = []string{"/var/www/mkd.sh", "/var/www/init.sh"}
)

const (
	dockerDir        = "/var/www"
	localhostDir     = "/Users/tsengyenchi/victor/php/tmp"
	gitRepo          = "git@github.com:nvizero/demo.git"
	dockerHubImg     = "19840112/firstphp:0.0.1"
	ControllerDir    = "/app/Http/Controllers/Admin/"
	ModelDir         = "/app/Models/"
	DatabaseDir      = "/database/migrations/"
	migration_table  = "docker exec %s bash -c \"cd /var/www && php artisan make:migration %s\""
	migration_table2 = "docker exec %s bash -c \"cd /var/www && php artisan make:migration create_%s_table --create=%s\""
	gen_model        = "docker exec %s bash -c \"cd /var/www && php artisan make:model %s\""
	docker_run_sh    = "docker exec %s bash %s"
	laravel_update   = "docker exec %s bash -c \"cd /var/www && /usr/local/bin/composer update\""
	CreateTableName  = "create_qoos"
	git_clone        = "cd %s && git clone %s %s "
	git_branch       = "cd %s && git branch %s"
	git_checkout     = "cd %s && git checkout %s"
	git_add          = "cd %s && git add ."
	git_commit       = "cd %s && git commit -m \"push branch %s\""
	git_push         = "cd %s && git push -u origin %s"
	gen_docker_cmd   = "docker run -d -p %s --name %s -v %s %s "
	ControllerName   = "Qoo"
	ModelName        = "Models/Qoo"
	TableName        = "qoos"
)
