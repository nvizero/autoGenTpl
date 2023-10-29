package control

var model_1 = `<?php
namespace App\Models;
use Kyslik\ColumnSortable\Sortable;
`

// class Post extends BaseModel
var model_2 = `
{
    /**
     * The attributes that are mass assignable.
     *	
     * @var array
     */
    use Sortable;
  `

// protected $table = 'posts';
var model_3 = `
    protected $fillable = [
  `

// 'title', 'category_id', 'content', 'sort', 'user_id', 'username' ,'img'
var model_4 = `
    ];
`

var model_5 = `
    public function tableFieldsSetting()
    {
        return  [
`

var model_6 = `       ];
    }
}
`
