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
        f.extractall(save_path)


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

def extract_4(args, tarobj):
    # ruleid: tarfile-extractall-traversal
    tf = tarfile.open(mode='r', fileobj=None)
    tf.extractall(save_path)

def extract_5(args, tarobj):
    # ok: tarfile-extractall-traversal
    tf = tarfile.open(mode='r', fileobj=None)
    tf.extractall(save_path, members=safu_members())

def extract_6(args, tarobj):
    # ok: tarfile-extractall-traversal
    tf = tarfile.open(mode='r', fileobj=None)
    tf.extractall(members=safu_members(), numeric_owner=True, path=save_path)

def extract_7(args, tarobj):
    # todoruleid: tarfile-extractall-traversal
    tf = tarfile.open(mode='r', fileobj=None)
    tf.extractall(members=None, numeric_owner=True, path=save_path)

def run():
    parser = get_parser()
    args = parser.parse_args()
    extract_1(args)
    extract_2(args)
    extract_3(args)
    extract_4(args)
    extract_5(args)
    extract_6(args)
    extract_7(args)


if __name__ == '__main__':
    run()
