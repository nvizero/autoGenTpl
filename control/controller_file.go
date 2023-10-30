package control

var controller_0 = `<?php
namespace App\Http\Controllers\Admin;
use Illuminate\Http\Request;
use App\Services\RequestService;
`

//use App\Models\Post;
//class PostController extends TemplateController

// public string $main = 'posts';

//      $this->entity = $posts;
//    function __construct(Request $request, Post $posts, RequestService $requestService)
//    {

var controller_1 = `
        $this->request = $request;
        $this->fieldsSetting = $this->entity->tableFieldsSetting();
        $this->requestService = $requestService;
    }
}

`
