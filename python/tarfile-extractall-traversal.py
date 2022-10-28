import argparse
import tarfile

def get_parser():
    parser = argparse.ArgumentParser(description='Process zip files.')
    parser.add_argument('--path', type=str, required=True,
                        help='Path to zip file')
    parser.add_argument('--save-path', type=str, required=True,
                            help='Output path')
    return parser


def extract_1(args):
    path = args.path
    save_path = args.save_path
    # ruleid: tarfile-extractall-traversal
    with tarfile.open(path) as f:
        
        import os
        
        def is_within_directory(directory, target):
            
            abs_directory = os.path.abspath(directory)
            abs_target = os.path.abspath(target)
        
            prefix = os.path.commonprefix([abs_directory, abs_target])
            
            return prefix == abs_directory
        
        def safe_extract(tar, path=".", members=None, *, numeric_owner=False):
        
            for member in tar.getmembers():
                member_path = os.path.join(path, member.name)
                if not is_within_directory(path, member_path):
                    raise Exception("Attempted Path Traversal in Tar File")
        
            tar.extractall(path, members, numeric_owner=numeric_owner) 
            
        
        safe_extract(f, save_path)


def extract_2(args):
    path = args.path
    save_path = args.save_path
    # ruleid: tarfile-extractall-traversal
    tarfile.open(path).extractall(save_path)


def extract_3(args):
    path = args.path
    save_path = args.save_path
    # ruleid: tarfile-extractall-traversal
    tf = tarfile.open(path)
    tf.extractall(save_path)


def run():
    parser = get_parser()
    args = parser.parse_args()
    #extract_1(args)
    #extract_2(args)
    extract_3(args)


if __name__ == '__main__':
    run()